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

import { Typography } from '@mui/material';
import { FC } from 'react';
import { EdgeInfoProps } from '../index';
import { groupSpecialFormat } from '../utils';

const General: FC<EdgeInfoProps> = ({ sourceName, sourceType, targetName }) => {
    return (
        <>
            <Typography variant='body2'>
                {groupSpecialFormat(sourceType, sourceName)} the ability to synchronize the password set by Local
                Administrator Password Solution (LAPS) on the computer {targetName}.
            </Typography>

            <Typography variant='body2'>
                The local administrator password for a computer managed by LAPS is stored in the confidential and
                Read-Only Domain Controller (RODC) filtered LDAP attribute{' '}
                <Typography component={'pre'}>ms-mcs-AdmPwd</Typography>.
            </Typography>
        </>
    );
};

export default General;
