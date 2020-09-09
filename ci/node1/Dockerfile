FROM node:alpine
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install
COPY node1.js ./
CMD ["node1.js"]