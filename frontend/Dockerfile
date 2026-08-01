# frontend/Dockerfile

# --- Build Stage ---
FROM node:24.18.0-alpine AS builder

WORKDIR /app

# Cache package installation
COPY package.json package-lock.json ./
RUN npm ci

# Copy the rest of the application files
COPY . .

# Build Nuxt (outputs to .output)
RUN npm run build

# --- Production Stage ---
FROM node:24.18.0-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV HOST=0.0.0.0
ENV PORT=3000

# Copy compiled output from builder
COPY --from=builder /app/.output ./.output

EXPOSE 3000

CMD ["node", ".output/server/index.mjs"]