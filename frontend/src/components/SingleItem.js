import React from 'react';
import { Button } from '@mui/material';

const SingleItem = ({ isActive }) => {
  return (
    <Button
      active={isActive ? 'true' : undefined}
    >
    </Button>
  );
};

export default SingleItem;
