import { useLocation } from 'react-router-dom';
import VoiceChat from './VoiceChat';

function VoiceCallWrapper() {
  const location = useLocation();
  
  // Pass location to the class-based component
  return <VoiceChat location={location} />;
}

export default VoiceCallWrapper;