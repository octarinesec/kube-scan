import React from "react";
import express from "express";
import { match, RouterContext } from "react-router";
import { renderToString } from "react-dom/server";
import fetch from "node-fetch";
import serialize from "serialize-javascript";
import runtimeConfig from "./config";

import getRoutes from "./routes";

const assets = require(process.env.RAZZLE_ASSETS_MANIFEST);
const path = require("path");

console.log(`> RAZZLE_PUBLIC_DIR: ${process.env.RAZZLE_PUBLIC_DIR}`);
console.log(`> PORT: ${process.env.PORT}`);
console.log(`> API_SERVER_HOST: ${process.env.API_SERVER_HOST}`);
console.log(`> API_SERVER_PORT: ${process.env.API_SERVER_PORT}`);

const host = process.env.API_SERVER_HOST || "http://localhost";
const port = process.env.API_SERVER_PORT || "8081";

const server = express();

server
  .disable("x-powered-by")
  .use(express.static(process.env.RAZZLE_PUBLIC_DIR));

const apiRouter = express.Router();
apiRouter.get("/", async function(req, res) {
  try {
    let url = `${host}:${port}${req.baseUrl.replace("/api", "")}`;
    console.log("Fetching ", url);
    const resultPromise = fetch(url);
    let fetchedData = await (await resultPromise).json();
    console.log("Fetched JSON:\n", JSON.stringify(fetchedData, null, 2));
    await res.json(fetchedData);
  } catch (err) {
    console.log("Error fetching API JSON:", err);
    await res.json({ ERROR: "no server connection" });
  }
});

apiRouter.post("/", async function(req, res) {
  try {
    let url = `${host}:${port}${req.baseUrl.replace("/api", "")}`;
    console.log("Requesting refresh ", url);
    await fetch(url, {method: 'post'})
    await res.json({ SUCCESS: "Refresh succeeded" });
  } catch (err) {
    console.log("Error fetching API JSON:", err);
    await res.json({ ERROR: "no server connection" });
  }
});

server.use("/api/*", apiRouter);

server.use((req, res) => {
  match(
    { routes: getRoutes(), location: req.url },
    (error, redirectLocation, renderProps) => {
      if (error) {
        res.status(500).send(error.message);
      } else if (redirectLocation) {
        res.redirect(302, redirectLocation.pathname + redirectLocation.search);
      } else if (renderProps) {
        // You can also check renderProps.components or renderProps.routes for
        // your "not found" component or route respectively, and send a 404 as
        // below, if you're using a catch-all route.
        const context = {};
        const markup = renderToString(<RouterContext {...renderProps} />);

        if (context.url) {
          res.redirect(context.url);
        } else {
          res.status(200).send(
            `<!doctype html>
                <html lang="">
                <head>
                    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
                    <meta charset="utf-8" />
                    <title>kube-scan</title>
                    <meta name="viewport" content="width=device-width, initial-scale=1">
                    <link rel="shortcut icon" href="favicon.ico" type="image/x-icon">
                    ${
                      assets.client.css
                        ? `<link rel="stylesheet" href="${assets.client.css}">`
                        : ""
                    }
                    ${
                      process.env.NODE_ENV === "production"
                        ? `<script src="${assets.client.js}" defer></script>`
                        : `<script src="${assets.client.js}" defer crossorigin></script>`
                    }
                </head>
                <body>
                    <div id="root">${markup}</div>
                    <div id="modal-root"></div>
                    <script>window.env = ${serialize(runtimeConfig)};</script>
                </body>
            </html>`
          );
        }
      } else {
        res.status(404).send("Not found");
      }
    }
  );
});

export default server;
