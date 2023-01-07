package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/pavinjat/assessment/expenses"
	"github.com/stretchr/testify/assert"
)

func TestGetAllExpensesHandler(t *testing.T) {
	seedExpense(t)
	var allexpenses []expenses.Expense

	res := request(http.MethodGet, uri("expenses"), nil)
	err := res.Decode(&allexpenses)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.Greater(t, len(allexpenses), 0)
}

func TestCreateExpenseHandler(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "TestObj",
		"amount": 100,
		"note": "Note for TestObj", 
		"tags": ["TestTag1", "TestTag2" ,"TestTag3", "TestTag4"]
	}`)
	var exp expenses.Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&exp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, exp.ID)
	assert.Equal(t, "TestObj", exp.Title)
	assert.Equal(t, 100, exp.Amount)
	assert.Equal(t, "Note for TestObj", exp.Note)
	/*assert.Equal(t, "TestTag1", exp.Tags)
	assert.Equal(t, "TestTag2", exp.Tags)*/
}

func TestGetExpenseHandler(t *testing.T) {
	c := seedExpense(t)

	var latest expenses.Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(c.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, c.ID, latest.ID)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Amount)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Tags)
}

func seedExpense(t *testing.T) expenses.Expense {
	var c expenses.Expense
	body := bytes.NewBufferString(`{
		"title": "TestObj",
		"amount": 100,
		"note": "Note for TestObj", 
		"tags": ["TestTag1", "TestTag2" ,"TestTag3", "TestTag4"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&c)
	if err != nil {
		t.Fatal("can't create uomer:", err)
	}
	return c
}

/*func TestUpdateExpenseHandler(t *testing.T) {
	t.Skip("TODO: implement me")
}*/

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
