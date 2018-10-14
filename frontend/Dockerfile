FROM node:8.12-alpine as builder

COPY package.json yarn.lock ./
RUN yarn install

COPY . .
RUN yarn prd-build

FROM nginx:mainline-alpine
COPY --from=builder /dist /usr/share/nginx/html
