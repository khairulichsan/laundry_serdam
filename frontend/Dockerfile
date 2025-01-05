# Gunakan Node.js image
FROM node:22

# Set working directory
WORKDIR /app

# Copy file ke container
COPY . .

# Install dependencies
RUN npm install

# Build aplikasi
RUN npm run build

# Gunakan HTTP server untuk menjalankan aplikasi
RUN npm install -g serve
CMD ["serve", "-s", "dist"]
