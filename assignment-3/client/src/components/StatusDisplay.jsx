import { Card, CardContent, Typography, Grid, Box } from "@mui/material";

function StatusDisplay({ status }) {
  const renderWaterWarning = () => {
    if (status.water <= 5) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="blue">
            Aman
          </Typography>
        </Box>
      );
    } else if (status.water >= 6 && status.water <= 8) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="yellow">
            Siaga
          </Typography>
        </Box>
      );
    } else if (status.water > 8) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="error">
            Bahaya
          </Typography>
        </Box>
      );
    }
  };

  const renderWindWarning = () => {
    if (status.wind <= 6) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="blue">
            Aman
          </Typography>
        </Box>
      );
    } else if (status.wind >= 7 && status.wind <= 15) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="yellow">
            Siaga
          </Typography>
        </Box>
      );
    } else if (status.wind > 15) {
      return (
        <Box mt={1}>
          <Typography variant="body2" color="error">
            Bahaya
          </Typography>
        </Box>
      );
    }
  };

  return (
    <Grid container justifyContent="center" spacing={2}>
      <Grid item>
        <Card>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Water
            </Typography>
            <Typography variant="h4">{status.water}</Typography>
            <Typography variant="h6">m</Typography>
            {renderWaterWarning()}
          </CardContent>
        </Card>
      </Grid>
      <Grid item>
        <Card>
          <CardContent>
            <Typography variant="h5" gutterBottom>
              Wind
            </Typography>
            <Typography variant="h4">{status.wind}</Typography>
            <Typography variant="h6">m/s</Typography>
            {renderWindWarning()}
          </CardContent>
        </Card>
      </Grid>
    </Grid>
  );
}

export default StatusDisplay;
