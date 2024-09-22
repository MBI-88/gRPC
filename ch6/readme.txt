Chapter 6 is for improving of some part of the chapter 5. We took it as tips.

Resume (tips):

Why is it not always more optimal to use varint for integer types?

R/ varint encoding serializes bigger numbers into a bigger amount of bytes

How can we get th number of bytes a message will be serialized into?

R/ using proto.Marshal + len

What kind of tag should we give  to fields that are often populated?

R/ Smaller tagas

What is the main problem of splitting messages to use smaller tags?

R/ We have overhead because sub-messages are serialized as length-delimited types

What is FileldMask?

R/ A collection of fields' paths telling us what data to include

When is a repeated field serialized as unpacked?

R/ When repeated fields are acting on complex types