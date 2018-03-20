#!/bin/sh

redis-server /redis.conf
echo "Redis server is running..."
redis-cache-for-lowgo