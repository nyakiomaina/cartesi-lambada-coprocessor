import React, { useEffect, useState } from 'react';
import { ethers } from 'ethers';
import { CID } from 'multiformats/cid';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Typography,
    TablePagination,
} from '@mui/material';
import { getProvider, getContract } from '../config';
const EventListener = () => {
    const [taskResponses, setTaskResponses] = useState([]);
    const [taskBatches, setTaskBatches] = useState([]); // For TaskBatchRegistered events
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);

    useEffect(() => {
        const setupListeners = async () => {
            const provider = getProvider();
            if (!provider) return;
            console.log('Provider:', provider.connection.url);

            const contract = getContract(provider);

            // Log contract address
            console.log('Contract address:', contract.address);

            // Log network info
            const network = await contract.provider.getNetwork();
            console.log('Connected to network:', network);

            // Listen to TaskResponded events
            contract.on('TaskResponded', (task, response) => {

                const { batchIndex, programId, inputHash } = task;
                const { resultCID, outputHash } = response;

                let decodedProgramId;
                let decodedResultCID;

                try {
                    decodedProgramId = CID.decode(Buffer.from(programId.slice(2), 'hex')).toString();
                    decodedResultCID = CID.decode(Buffer.from(resultCID.slice(2), 'hex')).toString();
                } catch (error) {
                    console.error('Error decoding CID:', error);
                    decodedProgramId = 'Invalid CID';
                    decodedResultCID = 'Invalid CID';
                }

                const newResponse = {
                    batchIndex: batchIndex.toNumber(),
                    programId: decodedProgramId,
                    inputHash: ethers.utils.hexlify(inputHash),
                    resultCID: decodedResultCID,
                    outputHash: ethers.utils.hexlify(outputHash),
                };
                setTaskResponses((prev) => [newResponse, ...prev]);
                console.log('TaskResponded:', newResponse);
            });

            // Listen to TaskBatchRegistered events
            contract.on(
            'TaskBatchRegistered',
            (index, blockNumber, merkleRoot, quorumNumbers, quorumThresholdPercentage, event) => {
                console.log('TaskBatchRegistered event received:', {
                    index,
                    blockNumber,
                    merkleRoot,
                    quorumNumbers,
                    quorumThresholdPercentage,
                });

                const newBatch = {
                    index: index.toNumber(),
                    blockNumber: blockNumber.toNumber(),
                    merkeRoot: ethers.utils.hexlify(merkleRoot),
                    quorumNumbers: ethers.utils.hexlify(quorumNumbers),
                    quorumThresholdPercentage: quorumThresholdPercentage.toNumber(),
                };
                setTaskBatches((prev) => [newBatch, ...prev]);
                console.log('TaskBatchRegistered:', newBatch);
            });

            // Cleanup
            return () => {
                contract.removeAllListeners('TaskResponded');
                contract.removeAllListeners('TaskBatchRegistered');
            };
        };

        setupListeners(); // Call the async function
    }, []);

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0);
    };

    return (
        <div>
            <Typography variant="h5" gutterBottom>
                Task Responded Events
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Batch Index</TableCell>
                            <TableCell>Program ID (CID)</TableCell>
                            <TableCell>Input Hash</TableCell>
                            <TableCell>Result CID</TableCell>
                            <TableCell>Output Hash</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {taskResponses
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((response, idx) => (
                                <TableRow key={idx}>
                                    <TableCell>{response.batchIndex}</TableCell>
                                    <TableCell>{response.programId}</TableCell>
                                    <TableCell>{response.inputHash}</TableCell>
                                    <TableCell>{response.resultCID}</TableCell>
                                    <TableCell>{response.outputHash}</TableCell>
                                </TableRow>
                            ))}
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    count={taskResponses.length}
                    page={page}
                    onPageChange={handleChangePage}
                    rowsPerPage={rowsPerPage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                    rowsPerPageOptions={[5, 10, 25]}
                />
            </TableContainer>

            <Typography variant="h5" gutterBottom style={{ marginTop: '20px' }}>
                Task Batch Registered Events
            </Typography>
            <TableContainer component={Paper}>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Batch Index</TableCell>
                            <TableCell>Block Number</TableCell>
                            <TableCell>Merke Root</TableCell>
                            <TableCell>Quorum Numbers</TableCell>
                            <TableCell>Quorum Threshold Percentage</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {taskBatches
                            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                            .map((batch, idx) => (
                                <TableRow key={idx}>
                                    <TableCell>{batch.index}</TableCell>
                                    <TableCell>{batch.blockNumber}</TableCell>
                                    <TableCell>{batch.merkeRoot}</TableCell>
                                    <TableCell>{batch.quorumNumbers}</TableCell>
                                    <TableCell>{batch.quorumThresholdPercentage}</TableCell>
                                </TableRow>
                            ))}
                    </TableBody>
                </Table>
                <TablePagination
                    component="div"
                    count={taskBatches.length}
                    page={page}
                    onPageChange={handleChangePage}
                    rowsPerPage={rowsPerPage}
                    onRowsPerPageChange={handleChangeRowsPerPage}
                    rowsPerPageOptions={[5, 10, 25]}
                />
            </TableContainer>
        </div>
    );
};

export default EventListener;
