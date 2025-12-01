#!/bin/bash
# rebuild.sh - Docker compose'u temizle ve yeniden başlat

echo "🛑 Stopping containers..."
docker-compose down

echo "🗑️  Removing old images..."
docker-compose rm -f
docker rmi driver-service 2>/dev/null || true

echo "🔨 Building fresh images..."
docker-compose build --no-cache

echo "🚀 Starting services..."
docker-compose up -d

echo "✅ Done! Checking status..."
docker-compose ps

echo ""
echo "📋 To view logs, run: docker-compose logs -f driver-service"