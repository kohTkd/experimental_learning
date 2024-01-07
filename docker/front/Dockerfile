FROM node:18.17-alpine

WORKDIR /app/

COPY ./package.json ./
COPY ./pnpm-lock.yaml ./
RUN npm i -g pnpm
RUN pnpm install

CMD ["pnpm", "run", "start"]
