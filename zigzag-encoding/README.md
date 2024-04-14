# ZigZag Encoding

## Overview
ZigZag encoding is a technique to reduce overall size of a number after encoding. The size will be based on the obsolute values as variable-length.

It's work by mapping negative values to positive values, and going back and forth (it's mean encode & decode from this positive number).

Example mapping:

|original value	| zigzag value |
| ----------------| -------------|
| -20	| 39
| -19	| 37
| -18	| 35
| -17	| 33
| -16	| 31
| -15	| 29
| -14	| 27
| -13	| 25
| -12	| 23
| -11	| 21
| -10	| 19

## Algorithm Description
### Zigzag encoding

```javascript
(i >> bitlength-1) ^ (i << 1)
```
with `i` being the number to be encoded, "^" being XOR-operation and ">>" would be arithmetic shifting-operation. Example for 32-bit integer, this would be:

`(i >> 31) ^ (i << 1)`

With `i = -1` 8 bits:

```javascript
// Step 1: i >> bitlength-1
11111111 = -1
11111111 = -1 >> 7

// Step 2: i << 1
11111111 = -1
11111110 = -1 << 1

// Step 3: (i >> bitlength-1) ^ (i << 1)
11111111 ^ 11111110 = 00000001 // = 1
```

With `i = 3` 8 bits:
```javascript
// Step 1: i >> bitlength-1
00000011 = 3
00000000 = 3 >> 7

// Step 2: i << 1
00000110 = 3 << 1

// Step 3: (i >> bitlength-1) ^ (i << 1)
00000000 ^ 00000110 = 00000110 // = 6
```

### Zigzad decoding
To convert a zigzag unsigned integer, i, we need to using the following formula below:
```javascript
(i >>> 1) ^ -(i & 1)
```
With zigzag-decoded value is `4` (8 bits, two's complement)
```javascript
// i = 4
// step 1: i >>> 1
00000100 = 4
00000010 = i >>> 1

// step 2: - (i & 1) (which equals ~(i & 1) + 1 => 2 complement)
00000100 & 00000001 = 00000101
~00000101 = 11111010 + 00000001 = 11111011

00000010 ^ 11111010 = 111111000 // = -4
```
