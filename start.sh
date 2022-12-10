#!/bin/sh

# SPDX-FileCopyrightText: 2021 The NGI Pointer Secure-Scuttlebutt Team of 2020/2021
#
# SPDX-License-Identifier: CC0-1.0

# we use absolute paths so this file can run easily within a docker container
/app/cmd/server/server -https-domain="${HTTPS_DOMAIN}" -repo="${REPO:-/ssb-go-room-secrets}" -aliases-as-subdomains="${ALIASES_AS_SUBDOMAINS}"
