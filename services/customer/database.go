// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package customer

import (
	"context"
	"errors"
	"math/rand"

	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"

	"github.com/puckpuck/hotrod/pkg/delay"
	"github.com/puckpuck/hotrod/pkg/log"
	"github.com/puckpuck/hotrod/pkg/tracing"
	"github.com/puckpuck/hotrod/services/config"
)

// database simulates Customer repository implemented on top of an SQL database
type database struct {
	tracer    opentracing.Tracer
	logger    log.Factory
	customers map[string]*Customer
	lock      *tracing.Mutex
}

func newDatabase(tracer opentracing.Tracer, logger log.Factory) *database {
	return &database{
		tracer: tracer,
		logger: logger,
		lock: &tracing.Mutex{
			SessionBaggageKey: "request",
		},
		customers: map[string]*Customer{
			"123": {
				ID:       "123",
				Name:     "Rachel's Floral Designs",
				Location: "115,277",
			},
			"567": {
				ID:       "567",
				Name:     "Amazing Coffee Roasters",
				Location: "211,653",
			},
			"392": {
				ID:       "392",
				Name:     "Trom Chocolatier",
				Location: "577,322",
			},
			"731": {
				ID:       "731",
				Name:     "Japanese Desserts",
				Location: "728,326",
			},
			"12323": {ID: "12323", Name: "Lambent Illumination", Location: "123,456"},
			"32392": {ID: "32392", Name: "Flux Water Gear", Location: "456,678"},
			"73451": {ID: "73451", Name: "Cipher Publishing", Location: "545,312"},
			"55673": {ID: "55673", Name: "Erudite Learning", Location: "546,853"},
			"44802": {ID: "44802", Name: "Quad Goals", Location: "545,385"},
			"18745": {ID: "18745", Name: "Obelus Concepts", Location: "321,583"},
			"23552": {ID: "23552", Name: "Zeal Wheels", Location: "879,556"},
			"23412": {ID: "23412", Name: "Moxie Marketing", Location: "455,687"},
			"23341": {ID: "23341", Name: "Bonefete Fun", Location: "231,159"},
			"69420": {ID: "69420", Name: "Bravura Inc", Location: "573,456"},
			"39001": {ID: "39001", Name: "Admire Arts", Location: "887,123"},
			"78945": {ID: "78945", Name: "Vortex Solar", Location: "321,456"},
			"59201": {ID: "59201", Name: "Sanguine Skincare", Location: "984,156"},
			"20885": {ID: "20885", Name: "Epic Adventure Inc", Location: "354,159"},
			"20482": {ID: "20482", Name: "Cogent Data", Location: "759,654"},
			"22083": {ID: "22083", Name: "Candor Corp", Location: "489,657"},
			"15864": {ID: "15864", Name: "Inspire Fitness Co", Location: "456,489"},
			"34320": {ID: "34320", Name: "Strat Security", Location: "654,357"},
			"38752": {ID: "38752", Name: "Innovation Arch", Location: "459,354"},
			"23155": {ID: "23155", Name: "Eco Focus", Location: "153,584"},
			"88654": {ID: "88654", Name: "Bicycles Green", Location: "784,215"},
			"35871": {ID: "35871", Name: "Farm Coffee", Location: "658,324"},
			"98457": {ID: "98457", Name: "Go Pastries", Location: "657,125"},
			"64853": {ID: "64853", Name: "Micro Planes", Location: "328,125"},
			"46321": {ID: "46321", Name: "Enhance Cars", Location: "954,236"},
			"57319": {ID: "57319", Name: "Value Walking", Location: "958,654"},
			"94317": {ID: "94317", Name: "Boost Fitness", Location: "326,125"},
			"96354": {ID: "96354", Name: "Vibrance Software", Location: "125,358"},
		},
	}
}

func (d *database) Get(ctx context.Context, customerID string) (*Customer, error) {
	d.logger.For(ctx).Info("Loading customer", zap.String("customer_id", customerID))

	// simulate opentracing instrumentation of an SQL query
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, "mysql")
		// #nosec
		span.SetTag("sql.query", "SELECT * FROM customer WHERE customer_id="+customerID)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}

	if !config.MySQLMutexDisabled {
		// simulate misconfigured connection pool that only gives one connection at a time
		d.lock.Lock(ctx)
		defer d.lock.Unlock()
	}

	// simulate RPC delay
	if (customerID == "57319" || customerID == "88654") && rand.Intn(100) < 50 {
		delay.Sleep(config.MySQLSlowCustomerDelay, config.MySQLSlowCustomerStdDev)
	} else {
		delay.Sleep(config.MySQLGetDelay, config.MySQLGetDelayStdDev)
	}

	if customer, ok := d.customers[customerID]; ok {
		return customer, nil
	}
	return nil, errors.New("invalid customer ID")
}
