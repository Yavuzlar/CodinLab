
import { hexToRGBA } from "@/utils/hex-to-rgba";
import { NorthRounded } from "@mui/icons-material";
import { Box, Button } from "@mui/material";
import { useEffect, useRef } from "react";

const ScrollTop = () => {
  const arrow = useRef();

  useEffect(() => {
    window.onscroll = () => {
      if (window.pageYOffset >= 200) {
        if (arrow.current) arrow.current.style.opacity = 1;
      } else {
        if (arrow.current) arrow.current.style.opacity = 0;
      }
    };
  }, []);

  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  return (
    <Box sx={{ overflow: 'hidden' }}>
      <Button
        ref={arrow}
        onClick={scrollToTop}
        aria-label="arrows"
        color="secondary"
        sx={{
          opacity: 0,
          position: 'fixed',
          transitionDuration: '500ms',
          transitionProperty: 'all',
          transitionTimingFunction: 'cubic-bezier(0.4, 0, 0.2, 1)',
          boxShadow: '0 0 5px rgba(0, 0, 0, 0.5)',
          borderRadius: '1.25rem',
          bottom: '1rem',
          right: '2rem',
          p: '0',
          zIndex: 99,
          color: '#eee',
          height: '40px',
          minWidth: '40px',
          width: '40px !important',
          backgroundColor: theme => `${hexToRGBA(theme.palette.background.paper, 1)} !important`,
          '&:hover': {
            transform: 'translate(0, -0.25rem) rotate(0) skewX(0) skewY(0) scaleX(1) scaleY(1)',
          }
        }}
      >
        <NorthRounded />
      </Button>
    </Box>
  );
};

export default ScrollTop;
