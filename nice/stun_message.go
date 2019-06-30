package nice

/**
 * StunClass:
 * @STUN_REQUEST: A STUN Request message
 * @STUN_INDICATION: A STUN indication message
 * @STUN_RESPONSE: A STUN Response message
 * @STUN_ERROR: A STUN Error message
 *
 * This enum is used to represent the class of
 * a STUN message, as defined in RFC5389
 */

/* Message classes */
type StunClass int
const (
	_ StunClass 	= iota
	STUN_REQUEST 	= 0
	STUN_INDICATION = 1
	STUN_RESPONSE	= 2
	STUN_ERROR		=3
)

/**
 * StunMethod:
 * @STUN_BINDING: The Binding method as defined by the RFC5389
 * @STUN_SHARED_SECRET: The Shared-Secret method as defined by the RFC3489
 * @STUN_ALLOCATE: The Allocate method as defined by the TURN draft 12
 * @STUN_SET_ACTIVE_DST: The Set-Active-Destination method as defined by
 * the TURN draft 4
 * @STUN_REFRESH: The Refresh method as defined by the TURN draft 12
 * @STUN_SEND: The Send method as defined by the TURN draft 00
 * @STUN_CONNECT: The Connect method as defined by the TURN draft 4
 * @STUN_OLD_SET_ACTIVE_DST: The older Set-Active-Destination method as
 * defined by the TURN draft 0
 * @STUN_IND_SEND: The Send method used in indication messages as defined
 * by the TURN draft 12
 * @STUN_IND_DATA: The Data method used in indication messages as defined
 * by the TURN draft 12
 * @STUN_IND_CONNECT_STATUS:  The Connect-Status method used in indication
 * messages as defined by the TURN draft 4
 * @STUN_CREATEPERMISSION: The CreatePermission method as defined by
 * the TURN draft 12
 * @STUN_CHANNELBIND: The ChannelBind method as defined by the TURN draft 12
 *
 * This enum is used to represent the method of
 * a STUN message, as defined by various RFCs
 */
/* Message methods */
type StunMethod int
const (
	_ StunMethod 		= iota
	STUN_BINDING		=	0x001    /* RFC5389 */
	STUN_SHARED_SECRET	=	0x002  /* old RFC3489 */
	STUN_ALLOCATE		=	0x003    /* TURN-12 */
	STUN_SET_ACTIVE_DST	=	0x004  /* TURN-04 */
	STUN_REFRESH		=	0x004  /* TURN-12 */
	STUN_SEND			=	0x004  /* TURN-00 */
	STUN_CONNECT		=	0x005    /* TURN-04 */
	STUN_OLD_SET_ACTIVE_DST=0x006  /* TURN-00 */
	STUN_IND_SEND		=	0x006    /* TURN-12 */
	STUN_IND_DATA		=	0x007    /* TURN-12 */
	STUN_IND_CONNECT_STATUS=0x008  /* TURN-04 */
	STUN_CREATEPERMISSION= 	0x008 /* TURN-12 */
	STUN_CHANNELBIND	= 	0x009 /* TURN-12 */
)

/**
 * StunAttribute:
 * @STUN_ATTRIBUTE_MAPPED_ADDRESS: The MAPPED-ADDRESS attribute as defined
 * by RFC5389
 * @STUN_ATTRIBUTE_RESPONSE_ADDRESS: The RESPONSE-ADDRESS attribute as defined
 * by RFC3489
 * @STUN_ATTRIBUTE_CHANGE_REQUEST: The CHANGE-REQUEST attribute as defined by
 * RFC3489
 * @STUN_ATTRIBUTE_SOURCE_ADDRESS: The SOURCE-ADDRESS attribute as defined by
 * RFC3489
 * @STUN_ATTRIBUTE_CHANGED_ADDRESS: The CHANGED-ADDRESS attribute as defined
 * by RFC3489
 * @STUN_ATTRIBUTE_USERNAME: The USERNAME attribute as defined by RFC5389
 * @STUN_ATTRIBUTE_PASSWORD: The PASSWORD attribute as defined by RFC3489
 * @STUN_ATTRIBUTE_MESSAGE_INTEGRITY: The MESSAGE-INTEGRITY attribute as defined
 * by RFC5389
 * @STUN_ATTRIBUTE_ERROR_CODE: The ERROR-CODE attribute as defined by RFC5389
 * @STUN_ATTRIBUTE_UNKNOWN_ATTRIBUTES: The UNKNOWN-ATTRIBUTES attribute as
 * defined by RFC5389
 * @STUN_ATTRIBUTE_REFLECTED_FROM: The REFLECTED-FROM attribute as defined
 * by RFC3489
 * @STUN_ATTRIBUTE_CHANNEL_NUMBER: The CHANNEL-NUMBER attribute as defined by
 * TURN draft 09 and 12
 * @STUN_ATTRIBUTE_LIFETIME: The LIFETIME attribute as defined by TURN
 * draft 04, 09 and 12
 * @STUN_ATTRIBUTE_MS_ALTERNATE_SERVER: The ALTERNATE-SERVER attribute as
 * defined by [MS-TURN]
 * @STUN_ATTRIBUTE_MAGIC_COOKIE: The MAGIC-COOKIE attribute as defined by
 * the rosenberg-midcom TURN draft 08
 * @STUN_ATTRIBUTE_BANDWIDTH: The BANDWIDTH attribute as defined by TURN draft 04
 * @STUN_ATTRIBUTE_DESTINATION_ADDRESS: The DESTINATION-ADDRESS attribute as
 * defined by the rosenberg-midcom TURN draft 08
 * @STUN_ATTRIBUTE_REMOTE_ADDRESS: The REMOTE-ADDRESS attribute as defined by
 * TURN draft 04
 * @STUN_ATTRIBUTE_PEER_ADDRESS: The PEER-ADDRESS attribute as defined by
 * TURN draft 09
 * @STUN_ATTRIBUTE_XOR_PEER_ADDRESS: The XOR-PEER-ADDRESS attribute as defined
 * by TURN draft 12
 * @STUN_ATTRIBUTE_DATA: The DATA attribute as defined by TURN draft 04,
 * 09 and 12
 * @STUN_ATTRIBUTE_REALM: The REALM attribute as defined by RFC5389
 * @STUN_ATTRIBUTE_NONCE: The NONCE attribute as defined by RFC5389
 * @STUN_ATTRIBUTE_RELAY_ADDRESS: The RELAY-ADDRESS attribute as defined by
 * TURN draft 04
 * @STUN_ATTRIBUTE_RELAYED_ADDRESS: The RELAYED-ADDRESS attribute as defined by
 * TURN draft 09
 * @STUN_ATTRIBUTE_XOR_RELAYED_ADDRESS: The XOR-RELAYED-ADDRESS attribute as
 * defined by TURN draft 12
 * @STUN_ATTRIBUTE_REQUESTED_ADDRESS_TYPE: The REQUESTED-ADDRESS-TYPE attribute
 * as defined by TURN-IPV6 draft 05
 * @STUN_ATTRIBUTE_REQUESTED_PORT_PROPS: The REQUESTED-PORT-PROPS attribute
 * as defined by TURN draft 04
 * @STUN_ATTRIBUTE_REQUESTED_PROPS: The REQUESTED-PROPS attribute as defined
 * by TURN draft 09
 * @STUN_ATTRIBUTE_EVEN_PORT: The EVEN-PORT attribute as defined by TURN draft 12
 * @STUN_ATTRIBUTE_REQUESTED_TRANSPORT: The REQUESTED-TRANSPORT attribute as
 * defined by TURN draft 12
 * @STUN_ATTRIBUTE_DONT_FRAGMENT: The DONT-FRAGMENT attribute as defined
 * by TURN draft 12
 * @STUN_ATTRIBUTE_XOR_MAPPED_ADDRESS: The XOR-MAPPED-ADDRESS attribute as
 * defined by RFC5389
 * @STUN_ATTRIBUTE_TIMER_VAL: The TIMER-VAL attribute as defined by TURN draft 04
 * @STUN_ATTRIBUTE_REQUESTED_IP: The REQUESTED-IP attribute as defined by
 * TURN draft 04
 * @STUN_ATTRIBUTE_RESERVATION_TOKEN: The RESERVATION-TOKEN attribute as defined
 * by TURN draft 09 and 12
 * @STUN_ATTRIBUTE_CONNECT_STAT: The CONNECT-STAT attribute as defined by TURN
 * draft 04
 * @STUN_ATTRIBUTE_PRIORITY: The PRIORITY attribute as defined by ICE draft 19
 * @STUN_ATTRIBUTE_USE_CANDIDATE: The USE-CANDIDATE attribute as defined by
 * ICE draft 19
 * @STUN_ATTRIBUTE_OPTIONS: The OPTIONS optional attribute as defined by
 * libjingle
 * @STUN_ATTRIBUTE_MS_VERSION: The MS-VERSION optional attribute as defined
 * by [MS-TURN]
 * @STUN_ATTRIBUTE_MS_XOR_MAPPED_ADDRESS: The XOR-MAPPED-ADDRESS optional
 * attribute as defined by [MS-TURN]
 * @STUN_ATTRIBUTE_SOFTWARE: The SOFTWARE optional attribute as defined by RFC5389
 * @STUN_ATTRIBUTE_ALTERNATE_SERVER: The ALTERNATE-SERVER optional attribute as
 * defined by RFC5389
 * @STUN_ATTRIBUTE_FINGERPRINT: The FINGERPRINT optional attribute as defined
 * by RFC5389
 * @STUN_ATTRIBUTE_ICE_CONTROLLED: The ICE-CONTROLLED optional attribute as
 * defined by ICE draft 19
 * @STUN_ATTRIBUTE_ICE_CONTROLLING: The ICE-CONTROLLING optional attribute as
 * defined by ICE draft 19
 * @STUN_ATTRIBUTE_MS_SEQUENCE_NUMBER: The MS-SEQUENCE NUMBER optional attribute
 * as defined by [MS-TURN]
 * @STUN_ATTRIBUTE_CANDIDATE_IDENTIFIER: The CANDIDATE-IDENTIFIER optional
 * attribute as defined by [MS-ICE2]
 * @STUN_ATTRIBUTE_MS_IMPLEMENTATION_VERSION: The IMPLEMENTATION-VERSION
 * optional attribute as defined by [MS-ICE2]
 * @STUN_ATTRIBUTE_NOMINATION: The NOMINATION attribute as defined by
 * draft-thatcher-ice-renomination-00 and deployed in Google Chrome
 *
 * Known STUN attribute types as defined by various RFCs and drafts
 */
/* Should be in sync with stun_is_unknown() */

type StunAttribute int
const (
	_ StunAttribute = iota
	STUN_ATTRIBUTE_MAPPED_ADDRESS	=	0x0001    /* RFC5389 */
	STUN_ATTRIBUTE_RESPONSE_ADDRESS	=	0x0002  /* old RFC3489 */
	STUN_ATTRIBUTE_CHANGE_REQUEST	=	0x0003    /* old RFC3489 */
	STUN_ATTRIBUTE_SOURCE_ADDRESS	=	0x0004    /* old RFC3489 */
	STUN_ATTRIBUTE_CHANGED_ADDRESS	=	0x0005  /* old RFC3489 */
	STUN_ATTRIBUTE_USERNAME			=	0x0006      /* RFC5389 */
	STUN_ATTRIBUTE_PASSWORD			=	0x0007    /* old RFC3489 */
	STUN_ATTRIBUTE_MESSAGE_INTEGRITY=	0x0008    /* RFC5389 */
	STUN_ATTRIBUTE_ERROR_CODE		=	0x0009      /* RFC5389 */
	STUN_ATTRIBUTE_UNKNOWN_ATTRIBUTES=	0x000A    /* RFC5389 */
	STUN_ATTRIBUTE_REFLECTED_FROM	=	0x000B    /* old RFC3489 */
	STUN_ATTRIBUTE_CHANNEL_NUMBER	=	0x000C        /* TURN-12 */
	STUN_ATTRIBUTE_LIFETIME			=	0x000D      /* TURN-12 */
	/* MS_ALTERNATE_SERVER is only used by Microsoft's dialect, probably should
	 * not to be placed in STUN_ALL_KNOWN_ATTRIBUTES */
	STUN_ATTRIBUTE_MS_ALTERNATE_SERVER=	0x000E /* MS-TURN */
	STUN_ATTRIBUTE_MAGIC_COOKIE		=	0x000F        /* midcom-TURN 08 */
	STUN_ATTRIBUTE_BANDWIDTH		=	0x0010      /* TURN-04 */
	STUN_ATTRIBUTE_DESTINATION_ADDRESS=	0x0011        /* midcom-TURN 08 */
	STUN_ATTRIBUTE_REMOTE_ADDRESS	=	0x0012    /* TURN-04 */
	STUN_ATTRIBUTE_PEER_ADDRESS		=	0x0012    /* TURN-09 */
	STUN_ATTRIBUTE_XOR_PEER_ADDRESS	=	0x0012    /* TURN-12 */
	STUN_ATTRIBUTE_DATA				=	0x0013     /* TURN-12 */
	STUN_ATTRIBUTE_REALM			=	0x0014      /* RFC5389 */
	STUN_ATTRIBUTE_NONCE			=	0x0015     /* RFC5389 */
	STUN_ATTRIBUTE_RELAY_ADDRESS	=	0x0016    /* TURN-04 */
	STUN_ATTRIBUTE_RELAYED_ADDRESS	=	0x0016    /* TURN-09 */
	STUN_ATTRIBUTE_XOR_RELAYED_ADDRESS=	0x0016    /* TURN-12 */
	STUN_ATTRIBUTE_REQUESTED_ADDRESS_TYPE=0x0017  /* TURN-IPv6-05 */
	STUN_ATTRIBUTE_REQUESTED_PORT_PROPS=0x0018  /* TURN-04 */
	STUN_ATTRIBUTE_REQUESTED_PROPS	=	0x0018  /* TURN-09 */
	STUN_ATTRIBUTE_EVEN_PORT		=	0x0018  /* TURN-12 */
	STUN_ATTRIBUTE_REQUESTED_TRANSPORT=	0x0019  /* TURN-12 */
	STUN_ATTRIBUTE_DONT_FRAGMENT	=	0x001A  /* TURN-12 */
	/* 0x001B */        /* reserved */
	/* 0x001C */        /* reserved */
	/* 0x001D */        /* reserved */
	/* 0x001E */        /* reserved */
	/* 0x001F */        /* reserved */
	STUN_ATTRIBUTE_XOR_MAPPED_ADDRESS=	0x0020    /* RFC5389 */
	STUN_ATTRIBUTE_TIMER_VAL		=	0x0021      /* TURN-04 */
	STUN_ATTRIBUTE_REQUESTED_IP		=	0x0022    /* TURN-04 */
	STUN_ATTRIBUTE_RESERVATION_TOKEN=	0x0022    /* TURN-09 */
	STUN_ATTRIBUTE_CONNECT_STAT		=	0x0023    /* TURN-04 */
	STUN_ATTRIBUTE_PRIORITY			=	0x0024     /* ICE-19 */
	STUN_ATTRIBUTE_USE_CANDIDATE	=	0x0025    /* ICE-19 */
	/* 0x0026 */        /* reserved */
	/* 0x0027 */        /* reserved */
	/* 0x0028 */        /* reserved */
	/* 0x0029 */        /* reserved */
	/* 0x002A-0x7fff */      /* reserved */

	/* Optional attributes */
	/* 0x8000-0x8021 */      /* reserved */
	STUN_ATTRIBUTE_OPTIONS			=	0x8001 /* libjingle */
	STUN_ATTRIBUTE_MS_VERSION		=	0x8008    /* MS-TURN */
	STUN_ATTRIBUTE_MS_XOR_MAPPED_ADDRESS=0x8020    /* MS-TURN */
	STUN_ATTRIBUTE_SOFTWARE			=	0x8022      /* RFC5389 */
	STUN_ATTRIBUTE_ALTERNATE_SERVER=	0x8023    /* RFC5389 */
	/* 0x8024 */        /* reserved */
	/* 0x8025 */        /* reserved */
	/* 0x8026 */        /* reserved */
	/* 0x8027 */        /* reserved */
	STUN_ATTRIBUTE_FINGERPRINT		=	0x8028    /* RFC5389 */
	STUN_ATTRIBUTE_ICE_CONTROLLED	=	0x8029    /* ICE-19 */
	STUN_ATTRIBUTE_ICE_CONTROLLING	=	0x802A    /* ICE-19 */
	/* 0x802B-0x804F */      /* reserved */
	STUN_ATTRIBUTE_MS_SEQUENCE_NUMBER=	0x8050     /* MS-TURN */
	/* 0x8051-0x8053 */      /* reserved */
	STUN_ATTRIBUTE_CANDIDATE_IDENTIFIER=0x8054     /* MS-ICE2 */
		/* 0x8055-0x806F */      /* reserved */
	STUN_ATTRIBUTE_MS_IMPLEMENTATION_VERSION=0x8070 /* MS-ICE2 */
		/* 0x8071-0xC000 */      /* reserved */
	STUN_ATTRIBUTE_NOMINATION=0xC001 /* https://tools.ietf.org/html/draft-thatcher-ice-renomination-00 */
		/* 0xC002-0xFFFF */      /* reserved */
)



/**
 * StunTransactionId:
 *
 * A type that holds a STUN transaction id.
 */
type StunTransactionId [STUN_MESSAGE_TRANS_ID_LEN]uint8
