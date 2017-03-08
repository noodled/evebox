/* Copyright (c) 2016 Jason Ish
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package server

import (
	"github.com/spf13/viper"
	"net/http"
)

type ConfigResponse struct {
	ElasticSearchIndex string                   `json:"ElasticSearchIndex"`
	EventServices      []map[string]interface{} `json:"event-services"`
	Extra              map[string]interface{}   `json:"extra"`
	Features           map[string]bool          `json:"features"`
}

func ConfigHandler(appContext AppContext, r *http.Request) interface{} {

	response := &ConfigResponse{}
	response.ElasticSearchIndex = viper.GetString("index")
	viper.UnmarshalKey("event-services", &response.EventServices)

	esKeyword := ""

	if appContext.ElasticSearch != nil {
		esKeyword, _ = appContext.ElasticSearch.GetKeywordType("")
	}

	// Make sure features is at least an empty list.
	response.Features = make(map[string]bool)

	for feature, enabled := range appContext.Features {
		response.Features[feature.String()] = enabled
	}

	response.Extra = map[string]interface{}{
		"elasticSearchKeyword": esKeyword,
	}

	return response
}
