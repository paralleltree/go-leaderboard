import 'normalize.css';

import * as React from 'react';
import { createRoot } from 'react-dom/client';

import { App } from './app';

const container = document.getElementById('root')
if (container === null) throw new Error('root element does not exists.');
const root = createRoot(container)
root.render(<App />);
