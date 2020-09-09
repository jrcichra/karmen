FROM node:13-alpine
WORKDIR /app
COPY node1.js package.json yarn.lock ./
RUN yarn install
CMD ["yarn start"]   