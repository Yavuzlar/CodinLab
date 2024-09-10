import {
  Box,
  Button,
  ListItemButton,
  ListItemIcon,
  ListItemText,
  Typography,
} from "@mui/material";
import { useRouter } from "next/router";
import { useNav } from "src/hooks/useNav";

const SingleItem = (props) => {
  const { title = "", path = null } = props;
  const { ChangePage } = useNav();

  const router = useRouter();
  const handleLocate = (p) => {
    return () => {
      if (p) {
        router.replace(p);
      }
      ChangePage();
    };
  };

  const style = {
    borderRadius: "50%",
    backgroundColor: "transparent",

    "&:hover": {
      backgroundColor: "rgba(255, 255, 255, 0.3)",
      borderRadius: "0.938rem",
    },
  };

  return (
    <Button
      sx={style}
      variant="text"
      onClick={handleLocate(path)}
      active={router.pathname == path}
    >
      <Typography
        sx={{
          textTransform: "capitalize",
          fontWeight: router.pathname === path ? 400 : 300,
          opacity: router.pathname === path ? 1 : 0.6,
          fontFamily: "Outfit",
        }}
      >
        {title}
      </Typography>
    </Button>
  );
};

export default SingleItem;
