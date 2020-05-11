# ----------------
# STEP 1:
# build with pkg
FROM node:12.16.2-alpine AS build
WORKDIR /app

# install dependencies with cache
COPY package.json .
COPY yarn.lock .
RUN yarn
# copy app files, build and package
COPY . /app

ENV PORT 8080
ENV PKG_CACHE_PATH .pkg-cache-upx

RUN yarn build --prod && npx pkg . -t node8-alpine-x64 --output app

# ----------------
# STEP 2:
# run with alpine
FROM alpine:latest
WORKDIR /app
ENV NODE_ENV=production

# install required libs
RUN apk update && apk add --no-cache libstdc++ libgcc

# copy prebuilt binary from previous step
COPY --from=build /app/app /app/app
COPY --from=build /app/build/public /app/build/public

CMD ["/app/app"]