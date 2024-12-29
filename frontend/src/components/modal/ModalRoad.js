import { Box, Button, Modal, Typography } from '@mui/material';
import React, { useState } from 'react'
import { theme } from 'src/configs/theme';

const ModalRoad = ({ buttonMessage = "", message = "" }) => {
    const [open, setOpen] = useState(false);
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    const modalStyle = {
        position: 'fixed',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        width: 500,
        bgcolor: 'background.paper',
        border: '2px solid aqua',
        borderRadius: "20px",
        boxShadow: 24,
        p: 4,
        zIndex: 1300,
    };


    return (
        <>
            <Button sx={{ backgroundColor: theme.palette.success.main, color: theme.palette.common.white }} onClick={handleOpen}>{buttonMessage}</Button>
            <Modal
                open={open}
                onClose={handleClose}
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
            >
                <Box sx={modalStyle}>
                    <Typography id="modal-modal-description" sx={{ mt: 2, whiteSpace: "pre-line" }}>
                        {message}
                    </Typography>
                </Box>
            </Modal>
        </>
    )
}

export default ModalRoad