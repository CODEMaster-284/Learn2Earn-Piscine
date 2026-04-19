#!/bin/bash
ls -1 | sed  '1d' | awk 'NR % 2 == 0'