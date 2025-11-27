Minimal Crypto Matching Engine

# What it does:

It takes buy/sell orders from users and matches them according to
preset rules

Limit order book is used:
It cosits of:
  - Bids -> buy orders
  - Asks -> sell orders
where each order contains:
  - Price
  - Amount
  - Timestamp
The order book is continuously changing as new orders enter and trades happen
Old orders are cancelled and prices move

# How orders get matched

  - Market orders:
    Orders to buy/sell immediatly at the best available price
  - Limit orders:
    Orders to buy/sell at specific price or better

# Matching rules

  - Best price first:
    Highest bidder gets priority
    Lowest seller gets priority
  - Earliest timestamp first
    If two users offer tht esam eprice, whoever submitted first gets priority

...

# What I will be doing

The aim of this project is to make a matching engine that clients connect to and send buy/sell requests that are then executed
This will be connected to via gRPC and currently it's not decided if there will be hosting or what the graphics of this will be


Ideally:
client -> gRPC -> matching engine -> output stream -> database -> candle stick graph -> frontend


Client (bot):
  Write data into a protobuf file (similar to json)
  -> .proto file gets serialised into golang code
  -> gets serialised into binary stream to enter server
Server:
  Obtains the protobuf data through the gRPC and does the internal matching stuff
Database:
  after the matching stuff this will enter the database (or an intermediate buffer if this access is too slow)


  # Usage

  If you have make installed you can just run

  `make`

  in the top level directory

  If you have docker instead

  `docker build -t app`

  `docker run --rm app`

# TODO: Finish updating this
