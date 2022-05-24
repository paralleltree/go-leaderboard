import * as React from 'react';
import { createRoot } from 'react-dom/client';

import { ApiClient } from './api_client';
import { SearchEvents } from './components/search_events';

const client = new ApiClient('http://localhost:8000');

const container = document.getElementById('root')
if (container === null) throw new Error('root element does not exists.');
const root = createRoot(container)
root.render(<SearchEvents client={client} />);
