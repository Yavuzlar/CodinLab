import { Card, CardContent, Typography } from "@mui/material";
import yavuzlarLogo from "../../../frontend/public/images/yavuzlar-logo-black.png";
import Image from "next/image";

import Link from "next/link";
import Translations from "src/components/Translations";

const Home = () => {
  return (
    <div>
      <Card>
        <CardContent>
          <Translations text={"home.title"} />
          <Translations text={"active_locale"} />
          <Translations text={"home.title"} />
          <Translations text={"home.content"} />
          {/* <Image src={yavuzlarLogo} /> */}
        </CardContent>
      </Card>
      <Card sx={{ mt: "10px" }}>
        <CardContent>
          <Translations text={"roads.title"} />
          <Translations text={"roads.content"} />
        </CardContent>
      </Card>
    </div>
  );
};

export default Home;
