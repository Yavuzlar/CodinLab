import { Card, CardContent, Typography } from "@mui/material";
import yavuzlarLogo from "../../../frontend/public/images/yavuzlar-logo-black.png";
import Image from "next/image";
const Home = () => {
  return (
    <div>
      <Card>
        <CardContent>
          <Typography>CodeInLab</Typography>
          <Image src={yavuzlarLogo} />
        </CardContent>
      </Card>
    </div>
  );
};

export default Home;
