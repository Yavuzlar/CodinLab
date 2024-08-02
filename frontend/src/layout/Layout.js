"use client";
import Navbar from "src/layout/components/Navbar";
import Footer from "src/layout/components/Footer";
import ScrollTop from "src/layout/components/ScrollTop";
import NextNProgress from "nextjs-progressbar";
import { useState, useEffect, Fragment } from "react";
import { Box, Container } from "@mui/material";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import { useRouter } from "next/router";
import navigation from "src/navigation";
import findParent from "src/utils/findParent";

const Layout = ({ children }) => {


  const router = useRouter();

  const [mounted, setMounted] = useState(false);
  const [navbarHeight, setNavbarHeight] = useState(88);
  const [footerHeight, setFooterHeight] = useState(290);

  // useEffect(() => {
  //     const navbar = document.getElementById("navbar");
  //     const footer = document.getElementById("footer");

  //     if (navbar) setNavbarHeight(navbar.offsetHeight);
  //     if (footer) setFooterHeight(footer.offsetHeight);
  // }, []);

  useEffect(() => {
    setMounted(true);
  }, []);

  const titles = findParent(navigation, router.pathname);

  if (!mounted) return <>{children}</>;

  return (
    <Fragment>
      <NextNProgress
        color="#3894d0"
        startPosition={0.3}
        stopDelayMs={200}
        height={3}
        showOnShallow={true}
        options={{ easing: "ease-in-out", speed: 500 }}
      />

      <Navbar />

      <Box sx={{ mt: "68px" }}>
        <Container maxWidth="lgPlus">
          <CustomBreadcrumbs titles={titles} />
        </Container>
      </Box>

      <Box
        sx={{
          width: "100%",
          minHeight: `calc(100vh - 192px)`,
          pb: "48px",
        }}
      >
        <Container maxWidth="lgPlus">
          <Box sx={{ mt: '27px' }}>
            {children}
          </Box>
        </Container>
      </Box>

      <Footer />
      {/* Other components */}
      <ScrollTop />
    </Fragment>
  );
};

export default Layout;
