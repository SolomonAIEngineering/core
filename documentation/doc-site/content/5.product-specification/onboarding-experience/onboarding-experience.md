# Onboarding Experience (Solomon AI)

Our goal is to make the onboarding process as smooth as possible on the Solomon AI platform. This document will describe the series of interactions involved in the onboarding process, considering both the client's perspective and the system's backend operations.

## Frontend
The onboarding process for our platform is a two-phased operation that integrates external authentication and internal user data collection, followed by a financial setup that leverages third-party services. Below is an in-depth technical walkthrough:

### 1. Initial Authentication Flow
- Prospective users initiate their journey by authenticating via Auth0, using OAuth2 providers like Google, Apple, Facebook, or Twitter.
- After successful authentication, Auth0 redirects the user to our platform's onboarding panel.

### 2. Platform-Specific Onboarding
- The onboarding panel serves as the entry point to the Solomon AI platform, where users must select their role: either as a solo-entrepreneur or a business entity.
- Based on the selection, a dynamic, user-centric form is displayed. This form, consisting of multiple steps, must be completed fully before proceeding.
- Upon form submission, a series of backend registration operations are executed, resulting in the generation of a new user record. The user ID of this record is then sent back to the frontend.

### 3. Data Aggregation and Billing Integration
- Following registration, a composite query gathers user data across various backend services to build a complete user profile.
- Users are then redirected to the billing section, powered by Stripe, which handles the financial transactions. Users complete their subscription through a Stripe payment link.
- Once Stripe confirms the payment, the user is redirected back to our platform.

### 4. Financial Account Integration
- The subsequent step involves linking the user's bank account, facilitated by Plaid. This integration is essential for the financial data synchronization in our platform.
- After the bank account connection is established, the user is navigated to the home dashboard.

### 5. Finalization and Sync Status
- The home dashboard provides the user with real-time updates on the bank account and transaction details synchronization process.


## Gateway 
The api gateway can only service requests upon sign up and authentication via the auth0 platform. Hence after onboading with auth0, all requests will be deemed "authorized" hence serviceable by the api gateway.


## Backend
### Account Creation

Upon the user's submission of the account creation request, the user service springs into action to process the request. This involves a multi-step operation where the financial service is called upon to execute a crucial role. Here, it's responsible for creating a financial record within the context of a distributed transaction, ensuring both consistency and reliability across the system's services.

Simultaneously, a Stripe record ID is generated, establishing a direct link with the user's financial data. This step is pivotal for facilitating seamless financial data integration and synchronization across the platform.

Additionally, an entry is created in Algolia, enhancing the searchability and accessibility of the user's information. This integration with Algolia not only speeds up data retrieval but also optimizes the overall search experience within the platform, enabling efficient and rapid access to user records.

### Stripe Payment & Webhook

When a user completes a payment via Stripe, our system initiates a streamlined process to acknowledge and record the transaction. A Stripe webhook is triggered, sending a notification to our application. This automated communication facilitates the creation of a subscription record, which is then meticulously stored in our datastore.

This process not only captures the transactional details but also meticulously categorizes the user's subscription level. By doing so, we maintain an accurate and up-to-date record of the different payment tiers and services that the user has elected to subscribe to, ensuring a cohesive and synchronized financial management system within our platform.

### Plaid Token Exchange And Account Sync
When a user grants us authorization to access their bank account, a token exchange process is initiated through Plaid. This allows us to convert the public link token into an access token. Using this access token, we will synchronize comprehensive details at the account level, including, but not limited to, account-specific information, transaction history, and more. Additionally, we will securely log all queries made during this process in our systems. It is important to note that we utilize AWS KMS (Key Management Service) to encrypt the access token, ensuring the highest level of security and data protection