"use client";
import Navbar from "@/layout/components/Navbar";
import Footer from "@/layout/components/Footer";
import ScrollTop from "@/layout/components/ScrollTop";
// import NextNProgress from "nextjs-progressbar";
import { useState, useEffect, Fragment } from "react";
import { Box, Container } from "@mui/material";
import CustomBreadcrumbs from "@/components/breadcrumbs";
import { useRouter } from "next/router";
import navigation from "@/navigation";
import findParent from "@/utils/findParent";

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


    const titles = findParent(navigation, router.pathname)

    if (!mounted) return <>{children}</>

    return (
        <Fragment>
            {/* <NextNProgress
                color="#e11d48"
                startPosition={0.3}
                stopDelayMs={200}
                height={3}
                showOnShallow={true}
                options={{ easing: "ease-in-out", speed: 500 }}
            /> */}
            {/* <Navbar /> */}

            <Box sx={{
                width: '100%', display: 'flex', py: '40px',
                minHeight: `calc(100vh - 80px)`,
                gap: '40px',
            }}>
                <Navbar />

                <Container maxWidth="xl">
                    <Box
                        sx={{
                            // minHeight: `calc(100vh - ${(navbarHeight || 0) + (footerHeight || 0)}px)`
                            position: 'relative',
                            minHeight: `calc(100%)`,
                            maxWidth: '100%',
                            width: '100%',
                            // ml: '200px'
                        }}
                    >
                        <CustomBreadcrumbs titles={titles} />

                        <Box sx={{ mt: '1rem' }}>
                            {children}
                        </Box>
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
