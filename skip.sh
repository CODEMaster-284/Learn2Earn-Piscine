#! /bin/bash
ls -1 | sed '1d' '1d; n; d'
