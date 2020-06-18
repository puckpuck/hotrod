/* k6 load script */

import http from "k6/http";
import {sleep} from "k6";

export let options = {
  stages: [
    {target: 3, duration: "0s"},
    {target: 3, duration: "30s"},
    {target: 5, duration: "20s"},
    {target: 7, duration: "10s"},
    {target: 7, duration: "20s"},
    {target: 5, duration: "10s"},
    {target: 5, duration: "5m"},
    {target: 3, duration: "30s"}
  ]
}

let lastRequestID = 0;
let customers = [123, 392, 731, 567];

export function setup() {
  return {clientUUID: Math.round(Math.random() * 10000)}
}

export default function(data) {
  lastRequestID++;
  let requestID = data.clientUUID + "-" + lastRequestID;
  let customer = customers[Math.floor(Math.random() * customers.length)];
  let url = 'http://localhost:8080/dispatch?customer=' + customer;
  let headers = {
      'jaeger-baggage': 'session=' + data.clientUUID + ', request=' + requestID
  };

  let res = http.get(url, {headers: headers});

  sleep(Math.random() * 2);
}