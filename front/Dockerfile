#FROM node:16-alpine AS base
#WORKDIR /app
#COPY package*.json ./
#
## ---- Dependencies ----
#FROM base AS dependencies
#RUN npm ci --only=production
#
## ---- Release ----
#FROM base AS release
#COPY --from=dependencies /app/node_modules ./node_modules
#COPY . .
#CMD npm start


FROM node:18-alpine as base
RUN apk add --no-cache g++ make py3-pip libc6-compat
WORKDIR /app
COPY package*.json ./
EXPOSE 3000

FROM base as builder
WORKDIR /app
COPY . .
RUN npm run build


FROM base as production
WORKDIR /app

ENV NODE_ENV=production
RUN npm ci

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
USER nextjs


COPY --from=builder --chown=nextjs:nodejs /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json
COPY --from=builder /app/public ./public

CMD npm start

FROM base as dev
ENV NODE_ENV=production
RUN npm install
COPY . .
CMD npm run dev
