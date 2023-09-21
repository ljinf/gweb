package main

import (
	"fmt"
	"github.com/ljinf/gweb"
)

func SubjectAddController(c *gweb.Context) error {
	c.SetOkStatus().Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *gweb.Context) error {
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *gweb.Context) error {
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *gweb.Context) error {
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *gweb.Context) error {
	subjectId, _ := c.ParamInt("id", 0)
	c.SetOkStatus().Json("ok, SubjectGetController:" + fmt.Sprint(subjectId))

	return nil
}

func SubjectNameController(c *gweb.Context) error {
	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}