//
// Copyright (c) 2011, Yanko D Sanchez Bolanos
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//     * Neither the name of the author nor the
//       names of its contributors may be used to endorse or promote products
//       derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package goprowl

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	API_SERVER = "https://api.prowlapp.com"
	ADD_PATH   = "/publicapi/add"
	API_URL    = API_SERVER + ADD_PATH
)

type Notification struct {
	Application string
	Description string
	Event       string
	Priority    string
	Providerkey string
	Url         string
}

type Goprowl struct {
	apikeys []string
}

type errorResponse struct {
	Error struct {
		Code    int    `xml:"code,attr"`
		Message string `xml:",chardata"`
	} `xml:"error"`
}

func (gp *Goprowl) RegisterKey(key string) error {

	if len(key) != 40 {
		return errors.New("Error, Apikey must be 40 characters long.")
	}

	gp.apikeys = append(gp.apikeys, key)
	return nil
}

func (gp *Goprowl) DelKey(key string) error {
	for i, value := range gp.apikeys {
		if strings.EqualFold(key, value) {
			copy(gp.apikeys[i:], gp.apikeys[i+1:])
			gp.apikeys[len(gp.apikeys) - 1] = ""
			gp.apikeys = gp.apikeys[:len(gp.apikeys) - 1]
			return nil
		}
	}
	return errors.New("Error, key not found")
}

func decodeError(def string, r io.Reader) (err error) {
	xres := errorResponse{}
	if xml.NewDecoder(r).Decode(&xres) != nil {
		err = errors.New(def)
	} else {
		err = errors.New(xres.Error.Message)
	}
	return
}

func (gp *Goprowl) Push(n *Notification) (err error) {

	keycsv := strings.Join(gp.apikeys, ",")

	vals := url.Values{
		"apikey":      []string{keycsv},
		"application": []string{n.Application},
		"description": []string{n.Description},
		"event":       []string{n.Event},
		"priority":    []string{n.Priority},
	}

	if n.Url != "" {
		vals["url"] = []string{n.Url}
	}

	if n.Providerkey != "" {
		vals["providerkey"] = []string{n.Providerkey}
	}

	r, err := http.PostForm(API_URL, vals)

	if err != nil {
		return
	} else {
		defer r.Body.Close()
		if r.StatusCode != 200 {
			err = decodeError(r.Status, r.Body)
		}
	}
	return
}
