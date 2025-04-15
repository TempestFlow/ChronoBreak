package conv

import (
	"errors"
	"strings"

	"github.com/Oudwins/zog"
)

func MapZogErr(errsMap map[string]zog.ZogIssue) error {
	var err error = nil
	if len(errsMap) > 0 {
		for field, issue := range errsMap {
			msg := []string{"Invalid:", field, ":", issue.String()}
			str := strings.Join(msg, " ")
			e := errors.New(str)
			if err == nil {
				err = e
			} else {
				err = errors.Join(err, e)
			}
		}
	}
	return err
}

func MapZogErrs(errsMap map[string][]*zog.ZogIssue) error {
	var err error = nil
	if len(errsMap) > 0 {
		for field, issues := range errsMap {
			for _, issue := range issues {
				msg := []string{"Invalid", field, ":", issue.Message}
				str := strings.Join(msg, " ")
				e := errors.New(str)
				if err == nil {
					err = e
				} else {
					err = errors.Join(err, e)
				}
			}
		}
	}
	return err
}
