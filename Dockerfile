FROM oven/bun:latest as builder

WORKDIR /app

COPY . .
RUN bun install
RUN bun build --target bun --format esm --outdir dist ./src/main.ts

FROM oven/bun:latest as runner

WORKDIR /app

COPY --from=builder /app/dist ./dist  
COPY --from=builder /app/src /app/
COPY --from=builder /app/node_modules ./node_modules

CMD ["bun", "dist/main.js"]