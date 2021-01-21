/* k6 load script */

import http from "k6/http";
import {sleep} from "k6";

export let options = {
  stages: [
    {target: 3, duration: "0s"},
    {target: 5, duration: "20s"},
    {target: 9, duration: "10s"},
    {target: 7, duration: "4m"},
    {target: 5, duration: "15m"},
    {target: 8, duration: "10s"},
    {target: 7, duration: "6m"},
    {target: 5, duration: "10m"},
    {target: 3, duration: "20s"}
  ]
}

const BASE_URL = "http://localhost:8080"
let lastRequestID = 0;
let customers = [12323, 32392, 73451, 55673, 44802, 18745, 23552, 
  23412, 23341, 69420, 39001, 78945, 59201, 20885, 
  20482, 22083, 15864, 34320, 38752, 23155, 88654,
  35871, 98457, 64853, 46321, 57319, 94317, 96354,
  57319, 57319, 88654, 88654];
let clients = {};

export default function() {
  lastRequestID++;
  let clientUUID = getClientUUID(__VU);
  let requestID = clientUUID + "-" + lastRequestID;
  let customer = customers[Math.floor(Math.random() * customers.length)];
  let url = BASE_URL + "/dispatch?customer=" + customer;
  let headers = {
      "jaeger-baggage": "session=" + clientUUID + ", request=" + requestID, 
      "client": clientUUID
  };

  let res = http.get(url, {headers: headers});

  sleep(Math.random() * 2);
}

function getClientUUID(vu) {
  if (!clients[vu]) {
    clients[vu] = Math.round(Math.random() * 10000000);
  }
  return clients[vu];
}