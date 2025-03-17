# PolySciChart Technology Stack

## Overview
This document defines the tech stack for PolySciChart, a set of programs to create and manage
PolySciChart posts (PSCPosts) and deliver them via X. The stack includes:
- chartgen - a PSCPost generator, previewer, and X poster.
- chartwriter - a PSCPost data writer; generates JSON and stores to S3 object.
- psc2x - a service that redirects polyscichart.com URL requests to corresponding  X post.

The stack is designed to be simple, low-cost, and easy to maintain. It uses a minimal set of
libraries and tools. The primary user traffic is to X and these tools are primarily management
and creation tools. Exception is the psc2x is a simple URL redirect service.

The stack prioritizes simplicity, low cost, and X integration for monetization. 

---

## Technology Stack

- Go language for 
- pure Go libarries
- S3 API for object data storage
- File system for local testing of data storage
- go-chart for chart generation
- go-twitter for X integration
- net/http for HTTP server
- OCI for object storage
- minio-go for S3 API support
- test and execute management programs on mac osx
- deploy psc2x to OCI VM using Docker conatainer image

