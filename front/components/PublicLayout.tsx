import { Fragment } from "react";

const PublicLayout = ({ children }: { children: React.ReactNode }) => {
    return (
        <Fragment>
            {children}
        </Fragment>
    );
};

export default PublicLayout;