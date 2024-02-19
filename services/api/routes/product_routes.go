package routes

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/Yasoobnaqvi/v4tech_task/v4_tech.git/utilities"
	"github.com/Yasoobnaqvi/v4tech_task/v4_tech.git/validators"
)



func Test(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()

	return c.Status(fiber.StatusCreated).JSON(response)
}

func AddProduct(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AddProductPostBody{}
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	} else {
		fireStore := c.Locals("firebaseFirestore").(*firestore.Client)
		if _, _, err := fireStore.Collection("users").Doc(c.Locals("UUID").(string)).
			Collection("categories").Doc(postBody.CategoryId).Collection("products").Add(c.Context(), postBody); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			delete(response, "status")
			response["message"] = "Product added successfully"
			return c.Status(fiber.StatusCreated).JSON(response)
		}
	}

}

func GetAllProducts(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	docId := c.Params("id")

	fireStore := c.Locals("firebaseFirestore").(*firestore.Client)
	if snapshot, err := fireStore.Collection("users").
		Doc(c.Locals("UUID").(string)).
		Collection("categories").Doc(docId).Collection("products").Documents(c.Context()).GetAll(); err != nil {
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		delete(response, "status")
		delete(response, "message")
		var products []validators.AddProductPostBody
		for _, item := range snapshot {
			var product validators.AddProductPostBody
			item.DataTo(&product)
			product.DocumentId = item.Ref.ID
			products = append(products, product)
		}
		response["data"] = &products
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func GetSingleProduct(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	cId := c.Params("cid")
	pId := c.Params("pid")
	firestore := c.Locals("firebaseFirestore").(*firestore.Client)
	if snapshot, err := firestore.Collection("users").Doc(c.Locals("UUID").(string)).
		Collection("categories").Doc(cId).Collection("products").Doc(pId).Get(c.Context()); err != nil {
		response["error"] = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	} else {
		delete(response, "message")
		delete(response, "status")
		response["data"] = snapshot.Data()
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func UpdateProduct(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AddProductPostBody{}
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		firestore := c.Locals("firebaseFirestore").(*firestore.Client)
		if _, err := firestore.Collection("users").Doc(c.Locals("UUID").(string)).
			Collection("categories").Doc(postBody.CategoryId).Collection("products").Doc(postBody.DocumentId).
			Set(c.Context(), postBody); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			delete(response, "status")
			response["message"] = "Record updated successfully"
			return c.Status(fiber.StatusNoContent).JSON(response)
		}
	}
}

func DeleteProduct(c *fiber.Ctx) error {
	response := utilities.GetBaseResponseObject()
	postBody := &validators.AddProductPostBody{}
	if err := utilities.PostBodyValidation(c, postBody); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	} else {
		firestore := c.Locals("firebaseFirestore").(*firestore.Client)
		if _, err := firestore.Collection("users").Doc(c.Locals("UUID").(string)).Collection("categories").Doc(postBody.CategoryId).
			Collection("products").Doc(postBody.DocumentId).
			Delete(c.Context()); err != nil {
			response["error"] = err.Error()
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		} else {
			delete(response, "status")
			response["message"] = "Record deleted successfully"
			return c.Status(fiber.StatusAccepted).JSON(response)
		}
	}
}