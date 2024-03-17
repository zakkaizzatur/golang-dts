import { CircularProgress } from "@mui/material";
import StatusDisplay from "../components/StatusDisplay";
import { useQuery } from "react-query";
import { getStatus } from "../services/statusService";
import { useEffect } from "react";

function StatusPage() {
  const {
    data: status,
    isLoading,
    isError,
    refetch,
  } = useQuery("status", getStatus, {
    staleTime: 15000, // Invalidate query every 15 seconds
  });

  useEffect(() => {
    const intervalId = setInterval(() => {
      refetch();
    }, 15000); // Refresh every 15 seconds

    return () => clearInterval(intervalId);
  }, [refetch]);

  if (isLoading) return <CircularProgress />;
  if (isError) return <div>Error fetching status data</div>;

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <div>
        <StatusDisplay status={status} />
      </div>
    </div>
  );
}

export default StatusPage;
