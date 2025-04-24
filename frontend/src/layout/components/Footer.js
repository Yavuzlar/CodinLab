import Image from 'next/image';
import Link from 'next/link';
import { Box, Container, Stack, IconButton, Typography } from '@mui/material';
import { GitHub, X, Instagram, LinkedIn, MenuBook } from '@mui/icons-material';
import { useTranslation } from "react-i18next";
import { styled } from '@mui/material/styles';
import Tooltip from '@mui/material/Tooltip';

// Styled components
const StyledFooter = styled(Box)(({ theme }) => ({
  backgroundColor: theme.palette.primary.dark,
  color: theme.palette.common.white,
  padding: theme.spacing(4, 0),
}));

const LogoWrapper = styled(Box)({
  position: 'relative',
  width: 160,
  height: 64,
});

const SocialButton = styled(IconButton)(({ theme }) => ({
  color: theme.palette.common.white,
  '&:hover': {
    backgroundColor: 'rgba(255, 255, 255, 0.1)',
  },
}));

const Footer = () => {

  const { t } = useTranslation();

  const socialLinks = [
    { icon: <MenuBook />, url: 'https://docs.yavuzlar.org/', label: 'Yavuzlar Docs' },
    { icon: <GitHub />, url: 'https://github.com/Yavuzlar', label: 'GitHub' },
    { icon: <X />, url: 'https://x.com/siberyavuzlar', label: 'X (Twitter)' },
    { icon: <LinkedIn />, url: 'https://www.linkedin.com/company/siberyavuzlar/', label: 'LinkedIn' },
    { icon: <Instagram />, url: 'https://www.instagram.com/siberyavuzlar/', label: 'Instagram' },
  ];

  return (
    <StyledFooter component="footer">
      <Container maxWidth="lg">
        <Stack
          direction={{ xs: "column", md: "row" }}
          justifyContent="space-between"
          alignItems="center"
          spacing={4}
        >
          {/* Logos */}
          <Stack direction="row" spacing={4} alignItems="center">
            <Link href="https://yavuzlar.org/" passHref target="_blank" rel="noopener noreferrer">
              <LogoWrapper>
                <Image
                  src="/images/yavuzlar-logo-text-team-white.png"
                  alt="Yavuzlar Logo"
                  fill
                  sizes="(max-width: 600px) 100vw, 160px"
                  style={{ objectFit: "contain" }}
                  priority
                />
              </LogoWrapper>
            </Link>
            <Link href="https://www.sibervatan.org/" passHref target="_blank" rel="noopener noreferrer">
              <LogoWrapper>
                <Image
                  src="/images/sibervatanlogo.svg"
                  alt="Siber Vatan Logo"
                  fill
                  sizes="(max-width: 600px) 100vw, 160px"
                  style={{ objectFit: "contain" }}
                  priority
                />
              </LogoWrapper>
            </Link>
          </Stack>

          {/* Social Links */}
          <Stack direction="row" spacing={2}>
            {socialLinks.map((link) => (
              <Tooltip title={link.label} key={link.label}>
                <SocialButton
                  component="a"
                  href={link.url}
                  target="_blank"
                  rel="noopener noreferrer"
                  aria-label={link.label}
                >
                  {link.icon}
                </SocialButton>
              </Tooltip>
            ))}
          </Stack>
        </Stack>

        {/* Copyright */}
        <Typography
          variant="body2"
          color="grey.400"
          textAlign={{ xs: "center", md: "left" }}
          sx={{ mt: 4 }}
        >
          © {new Date().getFullYear()} {t("footer.copyright")}
        </Typography>
      </Container>
    </StyledFooter>
  );
};

export default Footer;