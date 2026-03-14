
REDIS_CLI="redis-cli -h shared-redis -a 123456"

echo "Seeding Hot Ranking data"

$REDIS_CLI ZADD hot_ranking 21150 "b7f08794-0a5d-4b0b-a4a7-870c6103a860"
$REDIS_CLI ZADD hot_ranking 17300 "ed36c3ee-bf1c-4cc9-90f2-c6fee46d489a"
$REDIS_CLI ZADD hot_ranking 14800 "edd7a302-4ee4-4e35-b73c-110360967a9f"
$REDIS_CLI ZADD hot_ranking 12900 "887bf777-7618-4fed-9cd1-11a4094721f2"
$REDIS_CLI ZADD hot_ranking 11650 "53a624f3-3d3d-4e7b-9b15-4fae931d46a5"
$REDIS_CLI ZADD hot_ranking 9600  "9b06ca2a-e53f-4e6e-88ee-8b22229bac2a"
$REDIS_CLI ZADD hot_ranking 8500  "728800a1-e403-4e23-a321-1cdafe970e2f"
$REDIS_CLI ZADD hot_ranking 8400  "f19296cc-28b3-4777-8994-d1f9727d9fa9"
$REDIS_CLI ZADD hot_ranking 6950  "30c06be8-82a5-4a3e-a146-2cb230b5b10e"
$REDIS_CLI ZADD hot_ranking 5600  "21cef157-a9fa-4cbc-aa37-e81bc18a6ad2"
$REDIS_CLI ZADD hot_ranking 3700  "9b8470d0-3fbc-4552-910e-37fa0f198303"
$REDIS_CLI ZADD hot_ranking 2550  "f52c1a8c-6dc9-48f8-b954-d8f9991760eb"

echo ""
echo "Done"
$REDIS_CLI ZREVRANGE hot_ranking 0 -1 WITHSCORES
