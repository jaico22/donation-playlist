import react from "react";
import { useFbLogin } from "../FaceBookLoginProvider";

const Donation : React.FC = props => {
    const fbl = useFbLogin();
    return (
    <div>
        <b>Donating as {fbl?.fbUserId}</b>
    </div>);
}

export default Donation;