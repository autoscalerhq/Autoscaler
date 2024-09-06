"use client"

import React, { useCallback } from 'react';
import {
    ReactFlow,
    Controls,
    useNodesState,
    useEdgesState,
    addEdge,
    Node,
    Edge,
    OnConnectStartParams, Connection
} from '@xyflow/react';

import {ArrowDownUp, Bell, BookUp, ChevronsLeftRight, File, GitPullRequestDraft, SquareLibrary} from 'lucide-react';
import '@xyflow/react/dist/base.css';
import './index.css';
import TurboNode, { TurboNodeData } from './TurboNode';
import TurboEdge from './TurboEdge';
import FunctionIcon from './FunctionIcon';
import {EdgeBase} from "@xyflow/system";

const initialNodes: Node<TurboNodeData>[] = [
    {
        id: '0',
        position: { x: 0, y: 125 },
        data: { icon: <BookUp />, title: 'Push', subline: 'Service' },
        type: 'turbo',
    },
    {
        id: '1',
        position: { x: 0, y: 0 },
        data: { icon: <GitPullRequestDraft />, title: 'Pull Cloudwatch', subline: 'Aws' },
        type: 'turbo',
    },
    {
        id: '2',
        position: { x: 0, y: 250 },
        data: { icon: <ChevronsLeftRight />, title: 'Stream Kenisis', subline: 'Aws' },
        type: 'turbo',
    },
    {
        id: '3',
        position: { x: 500, y: 125 },
        data: { icon: <FunctionIcon />, title: 'Rules Engine', subline: 'Autoscaler' },
        type: 'turbo',
    },
    {
        id: '4',
        position: { x: 850, y: 0 },
        data: { icon: <ArrowDownUp />, title: 'Scaler', subline: 'Aws' },
        type: 'turbo',
    },
    {
        id: '5',
        position: { x: 1150, y: 125 },
        data: { icon: <Bell />, title: 'Notification ', subline: 'User' },
        type: 'turbo',
    },
    {
        id: '6',
        position: { x: 1500, y: 0 },
        data: { icon: <SquareLibrary />, title: 'Monitoring Connection', subline: 'Monitoring System' },
        type: 'turbo',
    },
];

const initialEdges: Edge[] = [
    {
        id: 'e0-3',
        source: '0',
        target: '3',
    },
    {
        id: 'e1-2',
        source: '1',
        target: '3',
    },
    {
        id: 'e3-4',
        source: '2',
        target: '3',
    },
    {
        id: 'e2-5',
        source: '3',
        target: '4',
    },
    {
        id: 'e4-5',
        source: '4',
        target: '5',
    },
    {
        id: 'e3-6',
        source: '3',
        target: '5',
    },
    {
        id: 'e5-6',
        source: '5',
        target: '6',
    },
    {
        id: 'e4-6',
        source: '4',
        target: '6',
    },
];

const nodeTypes = {
    turbo: TurboNode,
};

const edgeTypes = {
    turbo: TurboEdge,
};

const defaultEdgeOptions = {
    type: 'turbo',
    markerEnd: 'edge-circle',
};

const Flow = () => {
    const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
    const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);

    const onConnect = useCallback(
        (params: Connection | EdgeBase) => setEdges((els) => addEdge(params, els)),
        []
    );

    return (
        <ReactFlow
            nodes={nodes}
            edges={edges}
            onNodesChange={onNodesChange}
            onEdgesChange={onEdgesChange}
            onConnect={onConnect}
            fitView
            nodeTypes={nodeTypes}
            edgeTypes={edgeTypes}
            defaultEdgeOptions={defaultEdgeOptions}
        >
            <Controls showInteractive={false} />
            <svg>
                <defs>
                    <linearGradient id="edge-gradient">
                        <stop offset="0%" stopColor="#ae53ba" />
                        <stop offset="100%" stopColor="#2a8af6" />
                    </linearGradient>

                    <marker
                        id="edge-circle"
                        viewBox="-5 -5 10 10"
                        refX="0"
                        refY="0"
                        markerUnits="strokeWidth"
                        markerWidth="10"
                        markerHeight="10"
                        orient="auto"
                    >
                        <circle stroke="#2a8af6" strokeOpacity="0.75" r="2" cx="0" cy="0" />
                    </marker>
                </defs>
            </svg>
        </ReactFlow>
    );
};

export default Flow;
