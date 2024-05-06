import { Card, CardContent, Typography } from "@mui/material";
import yavuzlarLogo from "../../../frontend/public/images/yavuzlar-logo-black.png";
import Image from "next/image";
import { Button } from "@mui/material";

const Home = () => {
  return (
    <div>
      <Card>
        <CardContent>
          <Typography>CodeInLab</Typography>
          <Image src={yavuzlarLogo} />
          <Button> TEST </Button>
        </CardContent>
      </Card>
    </div>
  );
};

export default Home;
