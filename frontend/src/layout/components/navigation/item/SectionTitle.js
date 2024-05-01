import { Divider } from '@mui/material'
import { Fragment } from 'react'

const SectionTitle = ({ title = "" }) => {
    return (
        <Fragment>
            <Divider textAlign="left">
                {title}
            </Divider>
        </Fragment>
    )
}

export default SectionTitle