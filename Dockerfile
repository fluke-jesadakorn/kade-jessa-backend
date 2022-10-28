# 1 choose a compiler OS
FROM golang:alpine AS builder
# 2 (optional) label the compiler image
LABEL stage=builder
# 3 (optional) install any compiler-only dependencies
RUN apk add --no-cache gcc libc-dev
WORKDIR /workspace
# 4 copy all the source files
COPY . .
# 5 build the GO program
RUN CGO_ENABLED=0 GOOS=linux go build -a
# 6 choose a runtime OS
FROM alpine AS final

WORKDIR /
# 8 copy from builder the GO executable file
COPY --from=builder /workspace/kade-jessa .
# COPY --from=builder /workspace/_envs/env_$ENV.yaml ./_envs/
# 9 execute the program upon start 
ENV MONGODB_URI="mongodb+srv://kjimport:Ff_0813780670@cluster0.zo9yd.mongodb.net/?retryWrites=true&w=majority"
ENV MONGO_DATABASE="kade-jessa"
ENV MONGO_COLLECTION="products"
ENV GCS_BUCKET_NAME="kade-jessa"
ENV GIN_MODE="release"
ENV PORT=8080

EXPOSE 8080
CMD [ "./kade-jessa" ]