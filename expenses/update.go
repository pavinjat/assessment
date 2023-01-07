package expenses

import (
	"database/sql"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare statment for expense updated:" + err.Error()})
	}

	exp := Expense{}
	err = c.Bind(&exp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "The request could not be found:" + err.Error()})
	}

	_, err = stmt.Exec(id, exp.Title, exp.Amount, exp.Note, pq.Array(exp.Tags))
	intID, _ := strconv.Atoi(id)
	exp.ID = intID

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "Expense not found"})
	case nil:
		return c.JSON(http.StatusOK, exp)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "The system can't update:" + err.Error()})
	}

}
