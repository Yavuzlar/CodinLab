import React from 'react';
import Button from '@material-ui/core/Button';

const SingleItem = ({ isActive }) => {
  return (
    <Button
      active={isActive ? 'true' : undefined}
    >
      Item
    </Button>
  );
};

export default SingleItem;
