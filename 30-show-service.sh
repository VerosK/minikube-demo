#!/bin/bash

minikube service dummy --format '{{ .IP }}:{{ .Port }}' 



