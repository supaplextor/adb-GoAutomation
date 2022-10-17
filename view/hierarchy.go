// TODO : Documentation

package view

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/supaplextor/adbGoAutomation/geometry"
	"strconv"
	"strings"
)

type Hierarchy struct {
	XMLName  xml.Name `xml:"hierarchy"`     // Namespace of the hierarchy
	Rotation string   `xml:"rotation,attr"` // Rotation value of the hierarchy
	NodeList Nodes    `xml:"node"`          // Child nodes in the hierarchy
}

type Node struct {
	XMLName       xml.Name `xml:"node"`                // Namespace of the node
	Index         string   `xml:"index,attr"`          // Index of the node
	Resource      string   `xml:"resource-id,attr"`    // Index of the node
	Class         string   `xml:"class,attr"`          // Class of the node
	Package       string   `xml:"package,attr"`        // Package of the node
	Text          string   `xml:"text,attr"`           // Text of the node
	Description   string   `xml:"content-desc,attr"`   // Description of the node
	Checkable     string   `xml:"checkable,attr"`      // Checkable status of the node
	Checked       string   `xml:"checked,attr"`        // Checked status of the node
	Clickable     string   `xml:"clickable,attr"`      // Clickble status of the node
	Enabled       string   `xml:"enabled,attr"`        // Enabled status of the node
	Focusable     string   `xml:"focusable,attr"`      // Focusable status of the node
	Focused       string   `xml:"focused,attr"`        // Focused status of the node
	Scrollable    string   `xml:"scrollable,attr"`     // Scrollable status of the node
	LongClickable string   `xml:"long-clickable,attr"` // Long clickable status of the node
	Password      string   `xml:"password,attr"`       // Password field status of the node
	Selected      string   `xml:"selected,attr"`       // Selection status of the node
	Bounds        string   `xml:"bounds,attr"`         // View bounds of the node
	ChildNodes    Nodes    `xml:"node"`                // Child nodes for this node
}

type Nodes []Node

func (hierarchy Hierarchy) ConvertToViews() (Views, error) {
	return hierarchy.NodeList.ConvertToViews()
}

func (nodes Nodes) ConvertToViews() (Views, error) {
	views := Views{}
	for index := range nodes {
		v, err := nodes[index].ConvertToView()
		if err != nil {
			return Views{}, err
		}
		vv, err := nodes[index].ChildNodes.ConvertToViews()
		if err != nil {
			return Views{}, err
		}
		views = append(views, v)
		views = append(views, vv...)
	}
	return views, nil
}

func (node Node) ConvertToView() (View, error) {
	vw := View{}

	index, err := strconv.Atoi(node.Index)
	vw.Index = index
	if err != nil {
		return View{}, err
	}

	vw.Class = node.Class
	vw.Package = node.Package

	if strings.Contains(node.Resource, ":id/") {
		parts := strings.Split(node.Resource, ":id/")
		if len(parts) == 2 {
			vw.Resource = parts[1]
		} else {
			vw.Resource = node.Resource
		}
	} else {
		vw.Resource = node.Resource
	}

	vw.Text = node.Text
	vw.Description = node.Description

	vw.Clickable, _ = strconv.ParseBool(node.Clickable)
	vw.Checkable, _ = strconv.ParseBool(node.Checkable)
	vw.Checked, _ = strconv.ParseBool(node.Checked)
	vw.Enabled, _ = strconv.ParseBool(node.Enabled)
	vw.Focusable, _ = strconv.ParseBool(node.Focusable)
	vw.Focused, _ = strconv.ParseBool(node.Focused)
	vw.Scrollable, _ = strconv.ParseBool(node.Scrollable)
	vw.LongClickable, _ = strconv.ParseBool(node.LongClickable)
	vw.Password, _ = strconv.ParseBool(node.Password)
	vw.Selected, _ = strconv.ParseBool(node.Selected)

	bound_data := node.Bounds
	if strings.Contains(bound_data, "[") && strings.Contains(bound_data, "]") && strings.Contains(bound_data, ",") {
		parts := strings.Split(bound_data, "][")
		if len(parts) == 2 {
			top_left := parts[0]
			top_left = strings.Replace(top_left, "[", "", -1)

			bottom_right := parts[1]
			bottom_right = strings.Replace(bottom_right, "]", "", -1)

			top_parts := strings.Split(top_left, ",")
			bottom_parts := strings.Split(bottom_right, ",")

			if (len(top_parts) != 2) || (len(bottom_parts) != 2) {
				return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
			}

			top_x, err := strconv.Atoi(top_parts[0])
			if err != nil {
				return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
			}

			top_y, err := strconv.Atoi(top_parts[1])
			if err != nil {
				return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
			}

			bottom_x, err := strconv.Atoi(bottom_parts[0])
			if err != nil {
				return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
			}

			bottom_y, err := strconv.Atoi(bottom_parts[1])
			if err != nil {
				return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
			}

			vw.Bound = geometry.Rect{TopLeft: geometry.Point{top_x, top_y}, BottomRight: geometry.Point{bottom_x, bottom_y}}

			center_x := top_x + (bottom_x-top_x)/2
			center_y := top_y + (bottom_y-top_y)/2
			vw.Center = geometry.Point{X: center_x, Y: center_y}

		} else {
			return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
		}
	} else {
		return View{}, errors.New(fmt.Sprintf("Unable to determine bounds in [%v]", node))
	}

	return vw, nil
}
