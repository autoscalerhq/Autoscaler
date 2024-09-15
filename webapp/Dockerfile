# Stage 1: Install dependencies using Bun
FROM oven/bun:latest AS deps

WORKDIR /app

# Copy package and lock files
COPY package.json bun.lockb ./

# Install dependencies with Bun
RUN bun install

# Stage 2: Build the application
FROM oven/bun:latest AS builder

WORKDIR /app

# Copy the entire project
COPY . .

# Copy installed dependencies from the `deps` stage
COPY --from=deps /app/node_modules ./node_modules

# Build the application
RUN bun run build

# Stage 3: Prepare for production with standalone output
FROM node:20-alpine AS runner

WORKDIR /app

# Copy the built standalone output from the builder stage
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/static ./.next/static

# Expose the required port
EXPOSE 3333

# Set environment variables
ENV PORT 3333

# Run the built application using Next.js's standalone server
CMD ["node", "server.js"]