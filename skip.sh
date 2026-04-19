#!/bin/bash
ls -l | sed '1d' | awk 'NR % 2 == 0'
